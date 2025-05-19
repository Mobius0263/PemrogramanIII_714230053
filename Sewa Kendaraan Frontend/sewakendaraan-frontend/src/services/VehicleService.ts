import axios from 'axios';

const BASE_URL = 'http://localhost:3000/api/vehicles';

export const getVehicles = async () => {
  const res = await axios.get(BASE_URL);
  return res.data;
};

// You can also create postVehicle, updateVehicle, deleteVehicle, etc.
