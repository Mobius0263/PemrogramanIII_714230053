import React from 'react';
import { Link } from 'react-router-dom';

interface RentListItemProps {
  id: string;
  totalAmount: number;
}

const RentListItem: React.FC<RentListItemProps> = ({ id, totalAmount }) => {
  console.log("Rendering RentListItem with ID:", id); // Add this line for debugging
  return (
    <li>
      <Link to={`/rent/${id}`}>
        Rent ID: {id} - Amount: ${totalAmount}
      </Link>
    </li>
  );
};

export default RentListItem;