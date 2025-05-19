import React from 'react';

type ButtonProps = {
  children: React.ReactNode;
  onClick?: () => void;
};

const Button: React.FC<ButtonProps> = ({ children, onClick }) => (
  <button
    onClick={onClick}
    className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
  >
    {children}
  </button>
);

export default Button;
