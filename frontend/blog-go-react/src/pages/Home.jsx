import { useState, useEffect } from "react";

function Home() {
  const [posts, setPosts] = useState([]);

  const getPosts = async () => {
    try {
      const response = await fetch('http://localhost:3000/api/posts', {
        method: 'GET',
        credentials: 'include',
      });
      const result = await response.json();
      setPosts(result.data); 
    } catch (error) {
      console.log(error);
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
    </div>
  );
}

export default Home;
