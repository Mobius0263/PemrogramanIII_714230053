import React, { useEffect, useState } from 'react';
import axios from 'axios';

interface Rent {
  id: string;
  consumerId: string;
  consumerName: string;
  consumerPhone: string;
  vehicleId: string;
  vehicleBrand: string;
  vehicleModel: string;
  rentDate: string;
  returnDate: string;
  totalAmount: number;
}

interface Consumer {
  id: string;
  name: string;
  phoneNumber: string;
}

interface Vehicle {
  id: string;
  brand: string;
  model: string;
}

const RentListTemplate: React.FC = () => {
  const [rents, setRents] = useState<Rent[]>([]);
  const [consumers, setConsumers] = useState<Consumer[]>([]);
  const [vehicles, setVehicles] = useState<Vehicle[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [selectedRent, setSelectedRent] = useState<Rent | null>(null);
  const [formData, setFormData] = useState<Partial<Rent>>({});

  useEffect(() => {
    fetchRents();
    fetchConsumers();
    fetchVehicles();
  }, []);

  const fetchRents = async () => {
    try {
      const response = await axios.get<Rent[]>('http://localhost:3000/api/rents');
      setRents(response.data);
    } catch (error) {
      console.error('Error fetching rents:', error);
      setError('Failed to fetch rents');
    }
  };

  const fetchConsumers = async () => {
    try {
      const response = await axios.get<Consumer[]>('http://localhost:3000/api/consumers');
      setConsumers(response.data);
    } catch (error) {
      console.error('Error fetching consumers:', error);
      setError('Failed to fetch consumers');
    }
  };

  const fetchVehicles = async () => {
    try {
      const response = await axios.get<Vehicle[]>('http://localhost:3000/api/vehicles');
      setVehicles(response.data);
    } catch (error) {
      console.error('Error fetching vehicles:', error);
      setError('Failed to fetch vehicles');
    }
  };

  const createRent = async () => {
    try {
      const response = await axios.post<Rent>('http://localhost:3000/api/rents', formData);
      setRents([...rents, response.data]);
      setFormData({});
    } catch (error) {
      console.error('Error creating rent:', error);
      setError('Failed to create rent');
    }
  };

  const updateRent = async () => {
  if (!selectedRent) return;
  try {
    console.log('Updating rent:', selectedRent.id, formData);
    const correctId = selectedRent.id.replace(/^[0-9a-f]{12}([0-9a-f]{24})$/, '$1');
    console.log('Correct ID:', correctId);
    const response = await axios.put<Rent>(`http://localhost:3000/api/rents/${correctId}`, formData);
    console.log('Update response:', response.data);
    setRents(rents.map(rent => rent.id === selectedRent.id ? {...rent, ...response.data} : rent));
    setSelectedRent(null);
    setFormData({});
    fetchRents(); // Refresh the list after update
  } catch (error) {
    console.error('Error updating rent:', error);
    setError('Failed to update rent');
  }
};

  const deleteRent = async (id: string) => {
  if (window.confirm('Are you sure you want to delete this rent?')) {
    try {
      console.log('Deleting rent:', id);
      // Convert the ID to the correct format
      const correctId = id.replace(/^[0-9a-f]{12}([0-9a-f]{24})$/, '$1');
      await axios.delete(`http://localhost:3000/api/rents/${correctId}`);
      setRents(rents.filter(rent => rent.id !== id));
      fetchRents(); // Refresh the list after delete
    } catch (error) {
      console.error('Error deleting rent:', error);
      setError('Failed to delete rent');
    }
  }
};

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
  const { name, value } = e.target;
  
  switch (name) {
    case 'consumerId':
      const consumer = consumers.find(c => c.id === value);
      if (consumer) {
        setFormData(prev => ({
          ...prev,
          consumerId: consumer.id,
          consumerName: consumer.name,
          consumerPhone: consumer.phoneNumber
        }));
      }
      break;
    case 'vehicleId':
      const vehicle = vehicles.find(v => v.id === value);
      if (vehicle) {
        setFormData(prev => ({
          ...prev,
          vehicleId: vehicle.id,
          vehicleBrand: vehicle.brand,
          vehicleModel: vehicle.model
        }));
      }
      break;
    case 'totalAmount':
      const sanitizedValue = value.replace(/[^0-9.]/g, '');
      setFormData(prev => ({ ...prev, totalAmount: parseFloat(sanitizedValue) || 0 }));
      break;
    default:
      setFormData(prev => ({ ...prev, [name]: value }));
  }
};

  const validateForm = (): boolean => {
  if (!formData.consumerId || !formData.vehicleId ||
      !formData.rentDate || !formData.returnDate || 
      formData.totalAmount === undefined || formData.totalAmount === null) {
    setError('All fields are required');
    return false;
  }
  if (isNaN(Number(formData.totalAmount))) {
    setError('Total amount must be a valid number');
    return false;
  }
  setError(null);
  return true;
};

  const handleSubmit = async (e: React.FormEvent) => {
  e.preventDefault();
  if (validateForm()) {
    try {
      if (selectedRent) {
        await updateRent();
      } else {
        await createRent();
      }
      fetchRents(); // Refresh the list after create/update
    } catch (error) {
      console.error('Error submitting form:', error);
      setError('Failed to submit form');
    }
  }
};

  return (
    <div>
      <h2 className="text-2xl font-bold mb-4">Rent List</h2>
      {error && <div className="text-red-500 mb-4">{error}</div>}
      <form onSubmit={handleSubmit} className="mb-4">
        <select
          name="consumerId"
          value={formData.consumerId || ''}
          onChange={handleInputChange}
          className="mr-2 p-2 border"
        >
          <option value="">Select a consumer</option>
          {consumers.map(consumer => (
            <option key={consumer.id} value={consumer.id}>
              {consumer.name}
            </option>
          ))}
        </select>
        <input
          type="text"
          name="consumerPhone"
          value={formData.consumerPhone || ''}
          readOnly
          placeholder="Consumer Phone"
          className="mr-2 p-2 border bg-gray-100"
        />
        <select
          name="vehicleId"
          value={formData.vehicleId || ''}
          onChange={handleInputChange}
          className="mr-2 p-2 border"
        >
          <option value="">Select a vehicle</option>
          {vehicles.map(vehicle => (
            <option key={vehicle.id} value={vehicle.id}>
              {vehicle.brand} {vehicle.model}
            </option>
          ))}
        </select>
        <input
          type="text"
          name="vehicleBrand"
          value={formData.vehicleBrand || ''}
          readOnly
          placeholder="Vehicle Brand"
          className="mr-2 p-2 border bg-gray-100"
        />
        <input
          type="text"
          name="vehicleModel"
          value={formData.vehicleModel || ''}
          readOnly
          placeholder="Vehicle Model"
          className="mr-2 p-2 border bg-gray-100"
        />
        <input
          type="date"
          name="rentDate"
          value={formData.rentDate || ''}
          onChange={handleInputChange}
          className="mr-2 p-2 border"
        />
        <input
          type="date"
          name="returnDate"
          value={formData.returnDate || ''}
          onChange={handleInputChange}
          className="mr-2 p-2 border"
        />
        <input
          type="text"
          name="totalAmount"
          value={formData.totalAmount || ''}
          onChange={handleInputChange}
          placeholder="Total Amount"
          className="mr-2 p-2 border"
        />
        <button type="submit" className="p-2 bg-blue-500 text-white">
          {selectedRent ? 'Update Rent' : 'Create Rent'}
        </button>
      </form>
      <table className="min-w-full bg-white">
        <thead>
          <tr>
            <th className="py-2 px-4 border-b">ID</th>
            <th className="py-2 px-4 border-b">Consumer Name</th>
            <th className="py-2 px-4 border-b">Consumer Phone</th>
            <th className="py-2 px-4 border-b">Vehicle</th>
            <th className="py-2 px-4 border-b">Rent Date</th>
            <th className="py-2 px-4 border-b">Return Date</th>
            <th className="py-2 px-4 border-b">Total Amount</th>
            <th className="py-2 px-4 border-b">Actions</th>
          </tr>
        </thead>
        <tbody>
          {rents.map((rent) => (
            <tr key={rent.id}>
              <td className="py-2 px-4 border-b">{rent.id}</td>
              <td className="py-2 px-4 border-b">{rent.consumerName}</td>
              <td className="py-2 px-4 border-b">{rent.consumerPhone}</td>
              <td className="py-2 px-4 border-b">{`${rent.vehicleBrand} ${rent.vehicleModel}`}</td>
              <td className="py-2 px-4 border-b">{rent.rentDate}</td>
              <td className="py-2 px-4 border-b">{rent.returnDate}</td>
              <td className="py-2 px-4 border-b">${rent.totalAmount}</td>
              <td className="py-2 px-4 border-b">
                <button
                    onClick={() => {
                        setSelectedRent(rent);
                        setFormData({
                        ...rent,
                        consumerId: rent.consumerId,
                        vehicleId: rent.vehicleId,
                        totalAmount: rent.totalAmount
                        });
                    }}
                    className="mr-2 p-1 bg-yellow-500 text-white"
                    >
                    Edit
                    </button>
                    <button
                    onClick={() => deleteRent(rent.id)}
                    className="p-1 bg-red-500 text-white"
                    >
                    Delete
                    </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default RentListTemplate;