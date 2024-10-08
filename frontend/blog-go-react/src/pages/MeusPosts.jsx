import { useState, useEffect } from "react"

const MeusPosts = () => {

    /* 
    
    TODO: Fazer update e delete de posts
    
    */

  const [meusPosts, setMeusPosts] = useState([]); // state para armazenar os posts do usuario logado

  const getMeusPosts = async () => {
    try {
      const response = await fetch('http://localhost:3000/api/meusposts', { // dando fetch na api para pegar os posts do usuario logado
        method: 'GET',
        credentials: 'include',
      });
      const result = await response.json()
      console.log(result)
      if (result.length > 0) {
          setMeusPosts(result); // se o usuario tiver posts proprios, setamos os state com esses posts
      } else {
          setMeusPosts([])
      }
    } catch (error) {
      console.log(error);
    }
  }

    useEffect(() => {
        getMeusPosts(); // quando o componente MeusPosts for montado, faremos a chamada a api
    }, []);

  return (
    <div>
        {meusPosts.length > 0 ? ( // renderizacao condicional: se o usuario tiver posts, renderizamos os posts, se nao, renderizamos uma mensagem
            meusPosts.map((post) => ( // dando map no resultado da chamada a api: para cada post, renderizaremos um card com titulo, descricao, imagem e autor do post
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
            <p>Você não fez nenhum post ainda!</p>
        )}
    </div>
  )
}

export default MeusPosts