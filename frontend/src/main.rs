use yew::prelude::*;
//use local post crate

#[function_component]
fn App() -> Html {
    html! {
        <div>
            <h1>{"GetOSS"}</h1>
            <p>{"Forum for all developers and non-developers who want to keep their feet dry in the stormy ocean of Free & Open-sourced Software."}</p>
        </div>
    }
}

fn main() {
    yew::Renderer::<App>::new().render();
}
