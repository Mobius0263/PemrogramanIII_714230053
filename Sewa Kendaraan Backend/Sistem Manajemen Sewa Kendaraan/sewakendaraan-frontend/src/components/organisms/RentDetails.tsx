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

const RentDetails: React.FC = () => {
  const [rent, setRent] = useState<Rent | null>(null);
  const [error, setError] = useState<string | null>(null);
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();

  useEffect(() => {
    const fetchRent = async () => {
      if (!id) {
        setError('No rent ID provided');
        return;
      }

      try {
        const response = await axios.get(`http://localhost:3000/api/rents/${id}`);
        setRent(response.data);
      } catch (err) {
        if (axios.isAxiosError(err) && err.response?.status === 404) {
          setError('Rent not found');
        } else {
          setError('Failed to fetch rent details');
        }
        console.error(err);
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

export default RentDetails;