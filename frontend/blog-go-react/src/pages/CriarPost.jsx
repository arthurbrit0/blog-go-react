import { useState } from 'react';

function CriarPost() {
  const [titulo, setTitulo] = useState('');
  const [descricao, setDescricao] = useState('');
  const [imagem, setImagem] = useState(null);

  const handleSubmit = (e) => {
    e.preventDefault();
    // Handle post creation logic here
  };

  return (
    <div className="container mx-auto px-4 py-6 max-w-2xl">
      <h1 className="text-2xl font-bold mb-4">Create New Post</h1>
      <form onSubmit={handleSubmit} className="bg-white p-6 shadow-md rounded">
        <div className="mb-4">
          <label className="block text-gray-700">Title</label>
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
        <button type="submit" className="w-full bg-purple-500 text-white p-2 rounded">Criar Post</button>
      </form>
    </div>
  );
}

export default CriarPost;
