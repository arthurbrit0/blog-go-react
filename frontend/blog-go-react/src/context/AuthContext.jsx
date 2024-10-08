import { createContext, useContext, useState, useEffect } from 'react';

const AuthContext = createContext(); // Create the context

export const useAuth = () => {
  return useContext(AuthContext);
};

export const AuthProvider = ({ children }) => {
  const [auth, setAuth] = useState(null);
  const [loading, setLoading] = useState(true); // Add loading state

  const checkAuth = async () => {
    setLoading(true);
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
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    checkAuth();
  }, []);

  return (
    <AuthContext.Provider value={{ auth, checkAuth, loading }}>
      {children}
    </AuthContext.Provider>
  );
};
