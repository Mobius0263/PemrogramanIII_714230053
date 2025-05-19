import React from 'react';

interface DetailItemProps {
  label: string;
  value: string | number;
}

const DetailItem: React.FC<DetailItemProps> = ({ label, value }) => (
  <p>{label}: {value}</p>
);

export default DetailItem;