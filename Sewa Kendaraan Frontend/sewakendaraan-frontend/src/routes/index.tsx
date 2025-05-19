import { BrowserRouter, Routes, Route } from 'react-router-dom';
import HomePage from '../pages/HomePage.tsx';

const Router = () => (
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<HomePage />} />
    </Routes>
  </BrowserRouter>
);

export default Router;
