import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

function CriarPost() {
  const [titulo, setTitulo] = useState('');
  const [descricao, setDescricao] = useState('');
  const [imagem, setImagem] = useState(null);

  const navigate = useNavigate();

  const postar = async () => {
    const formData = new FormData();
    formData.append('titulo', titulo);
    formData.append('descricao', descricao);
    if (imagem) {
      formData.append('imagem', imagem);
    }
    try {
      const response = await fetch('http://localhost:3000/api/post', {
        method: 'POST',
        credentials: 'include',
        body: formData,
    });
    const data = await response.json();
    return data;
    } catch (error) {
      console.log(error)
    }
  }

  const handleSubmit = (e) => {
    e.preventDefault();
    postar();
    navigate('/meusposts');
  };

  return (
    <div className="container mx-auto px-4 py-6 max-w-2xl">
      <h1 className="text-2xl font-bold mb-4">Criar novo post</h1>
      <form onSubmit={handleSubmit} className="bg-white p-6 shadow-md rounded">
        <div className="mb-4">
          <label className="block text-gray-700">Título</label>
          <input 
            type="text" 
            className="w-full mt-1 p-2 border rounded" 
            value={titulo} 
            onChange={(e) => setTitulo(e.target.value)} 
            required 
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Descrição</label>
          <textarea 
            className="w-full mt-1 p-2 border rounded" 
            value={descricao} 
            onChange={(e) => setDescricao(e.target.value)} 
            required 
          ></textarea>
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Imagem</label>
          <input 
            type="file" 
            className="w-full mt-1 p-2" 
            onChange={(e) => setImagem(e.target.files[0])} 
          />
        </div>
        <button type="submit" className="w-full bg-blue-500 text-white p-2 rounded">Criar Post</button>
      </form>
    </div>
  );
}

export default CriarPost;
