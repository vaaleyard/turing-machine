use serde::Deserialize;

#[derive(Debug, Deserialize, Clone)]
pub enum Move {
    LEFT,
    RIGHT,
}

#[allow(dead_code)]
#[derive(Debug, Deserialize, Clone)]
#[serde(rename_all = "camelCase")]
pub struct Transition {
    pub simbolo: String,
    pub escreve: String,
    pub direcao: Move,
    pub estado_origem: String,
    pub estado_destino: String,
}

#[allow(dead_code)]
#[derive(Debug, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Machine {
    pub alfabeto: Vec<String>,
    simbolos: Vec<String>,
    pub transicoes: Vec<Transition>,
    pub marcador_inicio: String,
    pub estado_inicial: String,
    pub simbolo_branco: String,
    pub estados_finais: Vec<String>,
    estados: Vec<String>,
}
