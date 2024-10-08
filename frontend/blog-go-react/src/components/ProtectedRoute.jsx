
import { Navigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext.jsx'; 

const ProtectedRoute = ({ children }) => {
  const { auth, loading } = useAuth(); 

  if (loading) {
    return <div>Carregando...</div>; 
  }

  if (!auth) {
    return <Navigate to="/login" replace />;
  }

  return children;
};

export default ProtectedRoute;
