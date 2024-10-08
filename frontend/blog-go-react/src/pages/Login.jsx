import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Link } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

function Login() {
  const [email, setEmail] = useState(''); // criando states de email e senha para armazenar os valores dos inputs
  const [senha, setPassword] = useState('');

  const navigate = useNavigate(); // hook para navegar entre as rotas
  const { checkAuth } = useAuth(); // hook para verificar se o usuario esta autenticado

  const login = async (email, senha) => {
    try {
      const response = await fetch('http://localhost:3000/api/login', { // fazendo um post na rota de login com os dados dos inputs de email e senha
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include', // informndo que a aplicação vai usar cookies
        body: JSON.stringify({ email, senha }),
      });
      const data = await response.json();
      if (response.ok) {
        console.log(data); 
        checkAuth(); // checando o contexto de autenticacao para atualizar o estado do usuario autenticado
        navigate('/')
      }
      console.log(data);
    } catch (error) {
        console.error(error);
    }
    
  }

  const handleSubmit = (e) => {
    e.preventDefault();
    login(email, senha); // handleSubmit chama a função login passando os valores dos inputs após o submit do formulario de login
  };

  return (
    <div className="container mx-auto px-4 py-6 max-w-md">
      <h1 className="text-2xl font-bold mb-4">Login</h1>
      <form onSubmit={handleSubmit} className="bg-white p-6 shadow-md rounded">
        <div className="mb-4">
          <label className="block text-gray-700">Email</label>
          <input 
            type="email" 
            className="w-full mt-1 p-2 border rounded" 
            value={email} 
            onChange={(e) => setEmail(e.target.value)} // setando o valor do email como o valor do input
            required 
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Senha</label>
          <input 
            type="senha" 
            className="w-full mt-1 p-2 border rounded" 
            value={senha} 
            onChange={(e) => setPassword(e.target.value)} 
            required 
          />
        </div>
        <button type="submit" className="w-full bg-blue-500 text-white p-2 rounded">Login</button>
      </form>
      <Link to="/registrar" className="block text-center mt-4">Ainda não tem uma conta? Registre-se!</Link>
    </div>
  );
}

export default Login;
