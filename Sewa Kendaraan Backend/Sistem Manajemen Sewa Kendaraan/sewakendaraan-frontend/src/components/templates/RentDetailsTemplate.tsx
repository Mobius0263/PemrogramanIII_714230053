import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import axios from 'axios';

interface Rent {
  id: string;
  vehicleId: string;
  consumerId: string;
  rentDate: string;
  returnDate: string;
  totalAmount: number;
}

const RentDetailsTemplate: React.FC = () => {
  const [rent, setRent] = useState<Rent | null>(null);
  const [error, setError] = useState<string | null>(null);
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();

  useEffect(() => {
    console.log("Rent ID from params:", id);

    const fetchRent = async () => {
      if (!id) {
        setError('Invalid rent ID');
        return;
      }

      try {
        console.log(`Fetching rent details for ID: ${id}`);
        const response = await axios.get(`http://localhost:3000/api/rents/${id}`);
        console.log('Fetched rent details:', response.data);
        if (response.data && response.data.id) {
          setRent(response.data);
        } else {
          setError('Received invalid data from server');
        }
      } catch (err) {
        console.error('Error fetching rent details:', err);
        if (axios.isAxiosError(err)) {
          setError(`Failed to fetch rent details: ${err.response?.data?.error || err.message}`);
        } else {
          setError('An unexpected error occurred');
        }
      }
    };

    fetchRent();
  }, [id]);

  if (error) {
    return (
      <div>
        <h2>Error</h2>
        <p>{error}</p>
        <button onClick={() => navigate('/rents')}>Back to Rent List</button>
      </div>
    );
  }

  if (!rent) return <div>Loading...</div>;

  return (
    <div>
      <h2>Rent Details</h2>
      <p>ID: {rent.id}</p>
      <p>Vehicle ID: {rent.vehicleId}</p>
      <p>Consumer ID: {rent.consumerId}</p>
      <p>Rent Date: {rent.rentDate}</p>
      <p>Return Date: {rent.returnDate}</p>
      <p>Total Amount: ${rent.totalAmount}</p>
      <button onClick={() => navigate('/rents')}>Back to Rent List</button>
    </div>
  );
};

export default RentDetailsTemplate;