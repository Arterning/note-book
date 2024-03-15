

Create a store

```js
import { makeAutoObservable } from 'mobx';

class TodoStore {
    todos = [];

    constructor() {
        makeAutoObservable(this);
    }

// as you can see, we directly modified the todos instead of replace a new obj
// mobx 有意思的地方在于他违反了react的对象不可变性，可以直接修改todos对象
    addTodo = (text) => {
        this.todos.push({ id: Date.now(), text, done: false });
    };

    toggleTodo = (id) => {
        const todo = this.todos.find(todo => todo.id === id);
        if (todo) {
            todo.done = !todo.done;
        }
    };
}

export default new TodoStore();

```


Create a observer

```js
import { observer } from 'mobx-react-lite';
import todoStore from './store';

const TodoApp = observer(() => {
    const { todos, addTodo, toggleTodo } = todoStore;

    const handleAddTodo = () => {
        const text = prompt('Enter a new todo:');
        if (text) {
            addTodo(text);
        }
    };

    return (
        <div>
            <button onClick={handleAddTodo}>Add Todo</button>
            <ul>
                {todos.map(todo => (
                    <li key={todo.id} onClick={() => toggleTodo(todo.id)} style={{ textDecoration: todo.done ? 'line-through' : 'none' }}>
                        {todo.text}
                    </li>
                ))}
            </ul>
        </div>
    );
});

export default TodoApp;

```