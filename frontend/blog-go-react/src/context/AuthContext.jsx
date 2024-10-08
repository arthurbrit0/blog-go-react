import { createContext, useContext, useState, useEffect } from 'react';

const AuthContext = createContext(); // criando o contexto da autenticação para a aplicação toda

export const useAuth = () => {
  return useContext(AuthContext); // criando um hook para acessar o contexto da autenticação
};

export const AuthProvider = ({ children }) => {
  const [auth, setAuth] = useState(null); // criando o state que verifica se o usuario esta autenticado
  const [loading, setLoading] = useState(true); 

  const checkAuth = async () => {
    setLoading(true); // setando o loading como true para mostrar que a aplicação esta carregando
    try {
      const response = await fetch('http://localhost:3000/api/me', { // fazendo um get na rota /me para verificar se o usuario esta autenticado
        method: 'GET',                                              // a api so retorna as informacoes do usuario se ele estiver autenticado
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
