create table game
(
    id   text not null
        constraint game_pk
            primary key,
    name text not null,
    type text not null,
    data blob
);

create table betting_round
(
    id         text not null
        constraint betting_round_pk
            primary key,
    game_id    text not null
        references game,
    start_time TEXT not null,
    end_time   text not null,
    data       blob
);

create
unique index betting_round_id_uindex
    on betting_round (id);

create
unique index game_id_uindex
    on game (id);

create
unique index game_name_uindex
    on game (name);

create table user
(
    id            text not null
        constraint user_pk
            primary key,
    name          text not null,
    login         text not null,
    email         text,
    password      blob not null,
    password_salt blob not null
);

create table player
(
    id      text not null
        constraint player_pk
            primary key,
    user_id text not null
        references user,
    game_id text not null
        references game,
    stats   blob
);

create table bet
(
    id               text not null
        constraint bet_pk
            primary key,
    betting_round_id text not null
        references betting_round,
    game_id          text not null
        references game,
    player_id        text not null
        references player,
    data             blob
);

create
unique index bet_game_id_betting_round_id_player_id_uindex
    on bet (game_id, betting_round_id, player_id);

create
unique index bet_id_uindex
    on bet (id);

create
unique index player_id_uindex
    on player (id);

create
unique index player_user_id_game_id_uindex
    on player (user_id, game_id);

create
unique index user_id_uindex
    on user (id);

create
unique index user_login_uindex
    on user (login);

create
unique index user_name_uindex
    on user (name);
