create table pokemons(
    id serial primary key,
    name varchar,
    tag varchar,
    details varchar,
    img varchar
);

INSERT INTO pokemons(name, tag, details, img) VALUES
('Bulbasaur', '#001', 'Este animal tem uma planta nas costas e a usa para se defender', 'https://assets.pokemon.com/assets/cms2/img/pokedex/detail/001.png');