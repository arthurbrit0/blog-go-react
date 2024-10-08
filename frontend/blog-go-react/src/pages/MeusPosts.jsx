import { useState, useEffect } from "react";
import { Link } from "react-router-dom";

const MeusPosts = () => {
  const [meusPosts, setMeusPosts] = useState([]); 
  const [postParaEditar, setPostParaEditar] = useState(null); 
  const [tituloEditado, setTituloEditado] = useState(""); 
  const [descricaoEditada, setDescricaoEditada] = useState(""); 

  useEffect(() => {
    getMeusPosts();
  }, []);

  const getMeusPosts = async () => {
    try {
      const response = await fetch('http://localhost:3000/api/meusposts', { 
        method: 'GET',
        credentials: 'include',
      });
      const result = await response.json();
      if (result.length > 0) {
        setMeusPosts(result); 
      } else {
        setMeusPosts([]);
      }
    } catch (error) {
      console.log(error);
    }
  };

  const deletarPost = async (id) => {
    try {
      await fetch(`http://localhost:3000/api/posts/${id}`, { 
        method: 'DELETE',
        credentials: 'include',
      });
      setMeusPosts(meusPosts.filter(post => post.id !== id)); 
    } catch (error) {
      console.log(error);
    }
  };

  const abrirModalEditar = (post) => {
    setPostParaEditar(post); 
    setTituloEditado(post.titulo); 
    setDescricaoEditada(post.descricao);
  };

  const salvarEdicao = async () => {
    try {
      await fetch(`http://localhost:3000/api/posts/${postParaEditar.id}`, { // API de editar post
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
        body: JSON.stringify({
          titulo: tituloEditado,
          descricao: descricaoEditada,
        }),
      });
      setMeusPosts(meusPosts.map(post => post.id === postParaEditar.id ? { ...post, titulo: tituloEditado, descricao: descricaoEditada } : post));
      setPostParaEditar(null); 
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div>
      {meusPosts.length > 0 ? (
        meusPosts.map((post) => (
          <div key={post.id} className="mb-4 p-4 bg-white shadow-md rounded-lg">
            <h2 className="text-xl font-bold">{post.titulo}</h2>
            <p className="text-gray-700">{post.descricao}</p>
            <img src={post.imagem} alt={post.titulo} className="w-full h-64 object-cover mt-2" />
            <div className="mt-4">
              <h3 className="text-lg font-semibold">Autor: {post.usuario.primeiro_nome} {post.usuario.ultimo_nome}</h3>
            </div>
            <div className="mt-4 flex space-x-2">
              <button onClick={() => abrirModalEditar(post)} className="bg-blue-500 text-white p-2 rounded-lg">Editar</button>
              <button onClick={() => deletarPost(post.id)} className="bg-red-500 text-white p-2 rounded-lg">Deletar</button>
            </div>
          </div>
        ))
      ) : (
        <p>Você não fez nenhum post ainda!</p>
      )}

      {/* Modal de edição */}
      {postParaEditar && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center">
          <div className="bg-white p-4 rounded-lg shadow-lg max-w-lg w-full">
            <h2 className="text-2xl mb-4">Editar Post</h2>
            <input
              type="text"
              value={tituloEditado}
              onChange={(e) => setTituloEditado(e.target.value)}
              className="w-full p-2 border rounded mb-4"
              placeholder="Título"
            />
            <textarea
              value={descricaoEditada}
              onChange={(e) => setDescricaoEditada(e.target.value)}
              className="w-full p-2 border rounded mb-4"
              placeholder="Descrição"
            />
            <div className="flex justify-end space-x-2">
              <button onClick={() => setPostParaEditar(null)} className="bg-gray-300 p-2 rounded-lg">Cancelar</button>
              <button onClick={salvarEdicao} className="bg-green-500 text-white p-2 rounded-lg">Salvar</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default MeusPosts;
