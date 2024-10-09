import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Link } from 'react-router-dom';

function Registrar() {
  const [primeiro_nome, setFirstName] = useState(''); // criando states para armazenar os valores dos inputs
  const [ultimo_nome, setLastName] = useState('');
  const [email, setEmail] = useState('');
  const [senha, setPassword] = useState('');

  const navigate = useNavigate();   // usando o useNavigate para redirecionar o usuario apos o registro para a pagina de login

  const registrar = async (primeiro_nome, ultimo_nome, email, senha) => {
    try {
      const response = await fetch('http://localhost:3000/api/registrar', { // fazendo um post para a api com os dados informados pelo usuario
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        }, 
        body: JSON.stringify({ primeiro_nome, ultimo_nome, email, senha }), // passando os states, que serão os valores armazenados dos inputs, para o corpo da requisição
      });

      if (!response.ok) { // verificando se a resposta da api é ok
        const errorData = await response.json();
        throw new Error(errorData.mensagem || 'Erro ao registrar');
      }

      
      const data = await response.json();
      console.log(data); 
      navigate('/login') // se o registro for bem sucedido, o usuario sera redirecionado para a pagina de login
    } catch (error) {
      console.error('Erro ao registrar:', error);
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    registrar(primeiro_nome, ultimo_nome, email, senha); // handleSubmit sera chamado quando o formulario for submitado, registrando o usuario no banco de dados
  };

  return (
    <div className="container mx-auto px-4 py-6 max-w-md">
      <h1 className="text-2xl font-bold mb-4">Registrar</h1>
      <form onSubmit={handleSubmit} className="bg-white p-6 shadow-md rounded"> {/* AO SUBMITAR O FORMULARIO, A FUNCAO HANDLESUBMIT SERA CHAMADA, REGISTRANDO O USUARIO NO BANCO */}
        <div className="mb-4">
          <label className="block text-gray-700">Primeiro Nome</label>
          <input 
            type="text" 
            className="w-full mt-1 p-2 border rounded" 
            value={primeiro_nome} 
            onChange={(e) => setFirstName(e.target.value)}  // setando o valor do primeiro nome como o valor do input
            required 
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Último Nome</label>
          <input 
            type="text" 
            className="w-full mt-1 p-2 border rounded" 
            value={ultimo_nome} 
            onChange={(e) => setLastName(e.target.value)} // setando o valor do ultimo nome como o valor do input
            required 
          />
        </div>
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
            type="password" 
            className="w-full mt-1 p-2 border rounded" 
            value={senha} 
            onChange={(e) => setPassword(e.target.value)} // setando o valor da senha como o valor do input
            required 
          />
        </div>
        <button type="submit" className="w-full bg-green-500 text-white p-2 rounded">Registrar</button>
      </form>
      <Link to="/login" className="block text-center mt-4">Já tem uma conta? Faça login!</Link>
    </div>
  );
}

export default Registrar;
