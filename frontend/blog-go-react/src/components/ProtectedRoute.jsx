
import { Navigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext.jsx'; 

const ProtectedRoute = ({ children }) => {
  const { auth, loading } = useAuth();  // usando o hook customizado useAuth para verificar o estado de autenticacao do usuario

  if (loading) {
    return <div>Carregando...</div>; 
  }

  if (!auth) { // se o usuario não estiver autenticado, será redirecionado para a rota de login
    return <Navigate to="/login" replace />;
  }

  return children; // se o usuario estiver autenticado, os filhos do componente protegido pelo ProtectedRoute serão renderizados
};

export default ProtectedRoute;
