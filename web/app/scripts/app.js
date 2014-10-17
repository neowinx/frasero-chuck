(function(document) {
  'use strict';

  document.addEventListener('polymer-ready', function() {
    // Perform some behaviour
    console.log('Polymer is ready to rock!');

    document.querySelector('nueva-frase').addEventListener('frase-agregada', function(e) {
      var frasesElement = document.querySelector('frases-element');
      frasesElement.loadFrases();
    });

  });

// wrap document so it plays nice with other libraries
// http://www.polymer-project.org/platform/shadow-dom.html#wrappers
})(wrap(document));