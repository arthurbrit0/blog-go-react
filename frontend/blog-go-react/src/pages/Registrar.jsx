import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

function Registrar() {
  const [primeiro_nome, setFirstName] = useState('');
  const [ultimo_nome, setLastName] = useState('');
  const [email, setEmail] = useState('');
  const [senha, setPassword] = useState('');

  const navigate = useNavigate();   

  const registrar = async (primeiro_nome, ultimo_nome, email, senha) => {
    try {
      const response = await fetch('http://localhost:3000/api/registrar', {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        }, 
        body: JSON.stringify({ primeiro_nome, ultimo_nome, email, senha }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.mensagem || 'Erro ao registrar');
      }

      
      const data = await response.json();
      console.log(data); 
      navigate('/login')
    } catch (error) {
      console.error('Erro ao registrar:', error);
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    registrar(primeiro_nome, ultimo_nome, email, senha);
  };

  return (
    <div className="container mx-auto px-4 py-6 max-w-md">
      <h1 className="text-2xl font-bold mb-4">Registrar</h1>
      <form onSubmit={handleSubmit} className="bg-white p-6 shadow-md rounded">
        <div className="mb-4">
          <label className="block text-gray-700">Primeiro Nome</label>
          <input 
            type="text" 
            className="w-full mt-1 p-2 border rounded" 
            value={primeiro_nome} 
            onChange={(e) => setFirstName(e.target.value)} 
            required 
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Último Nome</label>
          <input 
            type="text" 
            className="w-full mt-1 p-2 border rounded" 
            value={ultimo_nome} 
            onChange={(e) => setLastName(e.target.value)} 
            required 
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Email</label>
          <input 
            type="email" 
            className="w-full mt-1 p-2 border rounded" 
            value={email} 
            onChange={(e) => setEmail(e.target.value)} 
            required 
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Senha</label>
          <input 
            type="password" 
            className="w-full mt-1 p-2 border rounded" 
            value={senha} 
            onChange={(e) => setPassword(e.target.value)} 
            required 
          />
        </div>
        <button type="submit" className="w-full bg-green-500 text-white p-2 rounded">Registrar</button>
      </form>
    </div>
  );
}

export default Registrar;
