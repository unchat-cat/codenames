import * as React from 'react';
import WordsPicker from '~/ui/words_picker';
import OriginalWords from '~/words.json';

// TODO: remove jquery dependency
// https://stackoverflow.com/questions/47968529/how-do-i-use-jquery-and-jquery-ui-with-parcel-bundler
var jquery = require('jquery');
window.$ = window.jQuery = jquery;

export const Lobby = ({ defaultGameID }) => {
  const [newGameName, setNewGameName] = React.useState(defaultGameID);
  const [selectedLanguage, setSelectedLanguage] = React.useState('Catalan');
  const [words, setWords] = React.useState(OriginalWords);

  function handleNewGame(e) {
    e.preventDefault();
    if (!newGameName) {
      return;
    }

    $.post(
      '/next-game',
      JSON.stringify({
        game_id: newGameName,
        word_set: words[selectedLanguage].split(', '),
	create_new: false,
      }),
      g => {
        const newURL = (document.location.pathname = '/' + newGameName);
        window.location = newURL;
      }
    );
  }

  return (
    <div id="lobby">
      <p id="banner">
        Traduït i adaptat per <a href="https://unchat.cat">unchat.cat</a>. Vegeu el <a href="https://github.com/unchat-cat/codenames">codi font</a>.
      </p>
      <div id="available-games">
        <form id="new-game">
          <p className="intro">
            Jugueu al codi secret en línia amb diversos dispositius en un tauler compartit.
	    Per crear un joc nou o afegir-vos a un joc existent, introduïu l'identificador
	    i feu click en 'Jugar!'.
          </p>
          <input
            type="text"
            id="game-name"
            autoFocus
            onChange={e => {
              setNewGameName(e.target.value);
            }}
            value={newGameName}
          />
          <button disabled={!newGameName.length} onClick={handleNewGame}>
            Jugar!
          </button>

          <div id="new-game-options">
            {Object.keys(OriginalWords).map(_language => (
              <WordsPicker
                key={_language}
                words={words[_language]}
                onWordChange={e => {
                  setWords({ ...words, [_language]: e.target.value });
                }}
                language={_language}
                selectedLanguage={selectedLanguage}
                onSelectedLanguageChange={() => {
                  setSelectedLanguage(_language);
                }}
              />
            ))}
          </div>
        </form>
      </div>
    </div>
  );
};
