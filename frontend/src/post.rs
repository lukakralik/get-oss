use yew::prelude::*;

enum Category {
    cli,
    ai,
    api,
    web,
    tool,
    game,
    embedded,
    vr,
}
pub struct Post {
    id: u32,
    name: String,
    author: Vec<String>,
    rating: f32,
    category: Category,
}

impl Post {
    fn new(&self) {
        todo!();
    }
}
