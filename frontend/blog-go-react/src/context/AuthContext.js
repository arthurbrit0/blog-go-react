import { createContext, useContext, useState, useEffect } from 'react';

const AuthContext = createContext(); // criamos o contexto da autenticação para a aplicação toda

export const useAuth = () => { // exportamos essa função, que é um hook personalizado para acessar o contexto de autenticação
  return useContext(AuthContext);
};

export const AuthProvider = ({ children }) => { // criamos o provedor de autenticação, que é um componente que envolve toda a aplicação
  const [auth, setAuth] = useState(null);    

  const checkAuth = async () => {
    try {
      const response = await fetch('http://localhost:3000/api/me', {
        method: 'GET',
        credentials: 'include',
      });

      if (response.ok) {
        const userData = await response.json();
        setAuth(userData);
      } else {
        setAuth(null); 
      }
    } catch (error) {
      setAuth(null);
    }
  };

  useEffect(() => {
    checkAuth();
  }, []);

  return (
    <AuthContext.Provider value={{ auth, checkAuth }}>
      {children}
    </AuthContext.Provider>
  );
};
