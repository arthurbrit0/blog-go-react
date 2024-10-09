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

      if (response.ok) { // se a resposta der ok, quer dizer que o usuario está autenticado
        const userData = await response.json();
        setAuth(userData); // autenticando o usuario e setando o state de auth como as informações do usuario
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
    checkAuth(); // toda vez que o componente for montado, a função checkAuth será chamada
  }, []);

  return (
    <AuthContext.Provider value={{ auth, checkAuth, loading }}> {/* passando o state de auth, a função checkAuth e o state de loading para o contexto */}
      {children}
    </AuthContext.Provider>
  );
};
