import { useState, useEffect } from "react";
import { Link } from "react-router-dom";

const MeusPosts = () => {
  const [meusPosts, setMeusPosts] = useState([]); // criando um array com os posts do usuario logado
  const [postParaEditar, setPostParaEditar] = useState(null); 
  const [tituloEditado, setTituloEditado] = useState(""); // criando states que serão editados no modal de edicao de cada post
  const [descricaoEditada, setDescricaoEditada] = useState(""); 

  useEffect(() => {
    getMeusPosts(); // a cada vez que o componente for montado, a funcao getMeusPosts sera chamada
  }, []);

  const getMeusPosts = async () => {
    try {
      const response = await fetch('http://localhost:3000/api/meusposts', {  // fazendo uma chamada a api, para pegar os posts do usuario logado
        method: 'GET',
        credentials: 'include',
      });
      const result = await response.json();
      if (result.length > 0) { // se houver algum post, setaremos os meusPosts como os dados retornados pela api
        setMeusPosts(result); 
      } else {
        setMeusPosts([]); // se não, apenas um array vazio
      }
    } catch (error) {
      console.log(error);
    }
  };

  const deletarPost = async (id) => { // nessa função, passaremos o id do post que queremos deletar. posteriormente, quando fizermos um map em cada post que pegamos, vincularemos
    try {                            // um botão para deletar com o id do post como parametro
      await fetch(`http://localhost:3000/api/posts/${id}`, {  // fazemos uma requisicao a api com o metodo delete e passando o id do post que queremos deletar
        method: 'DELETE',
        credentials: 'include',
      });
      setMeusPosts(meusPosts.filter(post => post.id !== id)); // depois que deletamos o post, setamos os meusPosts como um array com todos os posts, menos com o que acabou de ser deletado
    } catch (error) {
      console.log(error);
    }
  };

  const abrirModalEditar = (post) => { // recebemos um post como parâmetro para abrir o modal desse post especifico
    setPostParaEditar(post); // setamos o state do postParaEditar como o post que recebemos no parametro
    setTituloEditado(post.titulo);  // setamos o titulo do post para o titulo do post que recebemos no parametro
    setDescricaoEditada(post.descricao); // setamos a descricao do post para a descricao do post que recebemos no parametro
  };

  const salvarEdicao = async () => {
    try {
      await fetch(`http://localhost:3000/api/posts/${postParaEditar.id}`, { // fazemos uma requisicao para a api com o metodo put, passando o id do post que esta sendo editado 
        method: 'PUT',                                                      
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
        body: JSON.stringify({
          titulo: tituloEditado, // passando os states de tituloEditado e descricaoEditada para atualizar no banco de dados
          descricao: descricaoEditada,
        }),
      });
      /*
        depois que a edição é salva, atualizamos o array meusPosts para mapear em um novo array (no caso o mesmo meusPosts, mas com dados diferentes)
        os posts que não são o post que acabamos de editar, e conferimos se o post daquela iteração do map é igual ao post que estamos editando. se for,
        retornamos no map um novo objeto com os mesmos dados do post, mas com o titulo e descricao editados.   
      */
      setMeusPosts(meusPosts.map(post => post.id === postParaEditar.id ? { ...post, titulo: tituloEditado, descricao: descricaoEditada } : post));
      setPostParaEditar(null); // setamos o state postParaEditar como null
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div>
      {meusPosts.length > 0 ? ( // se o usuario tiver mais algum post, mostramos os posts desse usuario logado
        meusPosts.map((post) => (
          <div key={post.id} className="mb-4 p-4 bg-white shadow-md rounded-lg">
            <h2 className="text-xl font-bold">{post.titulo}</h2>
            <p className="text-gray-700">{post.descricao}</p>
            <img src={post.imagem} alt={post.titulo} className="w-full h-64 object-cover mt-2" />
            <div className="mt-4">
              <h3 className="text-lg font-semibold">Autor: {post.usuario.primeiro_nome} {post.usuario.ultimo_nome}</h3>
            </div>
            <div className="mt-4 flex space-x-2">
              <button onClick={() => abrirModalEditar(post)} className="bg-blue-500 text-white p-2 rounded-lg">
                Editar {/* Quando o usuário clica no botão de editar, a função abrirModal editar será chamda, tendo como parâmetro o post que está sendo renderizado no card*/}
              </button>
              <button onClick={() => deletarPost(post.id)} className="bg-red-500 text-white p-2 rounded-lg">
                Deletar {/* Quando o usuário clica no botão de deletar, fazemos uma chamada a api para deletar o post que está sendo renderizado no card*/}
              </button>
            </div>
          </div>
        ))
      ) : (
        <p>Você não fez nenhum post ainda!</p>
      )}

      {postParaEditar && ( // modal para editar o post, que so sera renderizado se algum post tiver sendo editado
        <div className="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center">
          <div className="bg-white p-4 rounded-lg shadow-lg max-w-lg w-full">
            <h2 className="text-2xl mb-4">Editar Post</h2>
            <input // os inputs atualizarão o valor dos estados tituloEditado e descricaoEditada com seus valores
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
              <button onClick={() => setPostParaEditar(null)} className="bg-gray-300 p-2 rounded-lg">
                Cancelar {/* Quando clicamos no botão para cancelar, o postParaEditar fica como null, o que faz o modal ser fechado */}
              </button>
              <button onClick={salvarEdicao} className="bg-green-500 text-white p-2 rounded-lg">
                Salvar {/* Quando clicamos no botão de salvar, a função salvarEdição é chamada, atualizando os dados do post com os dados passados pelo usuario nos inputs do modal */}  
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default MeusPosts;
