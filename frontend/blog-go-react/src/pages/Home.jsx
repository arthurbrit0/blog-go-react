import { useState, useEffect } from "react";

function Home() {
  const [posts, setPosts] = useState([]);
  const [paginaAtual, setPaginaAtual] = useState(1); 
  const [ultimaPagina, setUltimaPagina] = useState(1); 
  const [totalPosts, setTotalPosts] = useState(0);

  const getPosts = async (pagina = 1) => {
    try {
      const response = await fetch(`http://localhost:3000/api/posts?pagina=${pagina}`, {
        method: 'GET',
        credentials: 'include',
      });
      const result = await response.json();
      setPosts(result.data); 
      setPaginaAtual(pagina);
      setUltimaPagina(result.meta.ultima_pagina); 
      setTotalPosts(result.meta.total);
    } catch (error) {
      console.log(error);
    }
  };

    const paginaAnterior = () => {
      if (paginaAtual > 1) {
        getPosts(paginaAtual - 1);
      }
    };
  
    const proximaPagina = () => {
      if (paginaAtual < ultimaPagina) {
        getPosts(paginaAtual + 1);
      }
    };

  useEffect(() => {
    getPosts();
  }, []);

  return (
    <div>
      {posts.length > 0 ? (
        posts.map((post) => (
          <div key={post.id} className="mb-4 p-4 bg-white shadow-md rounded-lg">
            <h2 className="text-xl font-bold">{post.titulo}</h2>
            <p className="text-gray-700">{post.descricao}</p>
            <img src={post.imagem} alt={post.titulo} className="w-full h-64 object-cover mt-2" />
            <div className="mt-4">
              <h3 className="text-lg font-semibold">Autor: {post.usuario.primeiro_nome} {post.usuario.ultimo_nome}</h3>
            </div>
          </div>
        ))
      ) : (
        <p>Carregando posts...</p>
      )}

      {ultimaPagina > 1 && (
        <div className="flex justify-between mt-4">
          <button
            onClick={paginaAnterior}
            disabled={paginaAtual === 1}
            className="bg-gray-200 p-2 rounded-lg"
          >
            Anterior
          </button>
          <span>{`Página ${paginaAtual} de ${ultimaPagina}`}</span>
          <button
            onClick={proximaPagina}
            disabled={paginaAtual === ultimaPagina}
            className="bg-gray-200 p-2 rounded-lg"
          >
            Próxima
          </button>
        </div>
      )}
    </div>
  );
}

export default Home;
