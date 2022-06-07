var elixir = require('laravel-elixir');

require('laravel-elixir-webpack-official');

require('laravel-elixir-vue');

elixir(function(mix) {

    mix.webpack('vue-cookie.js', 'build', 'src');

});
