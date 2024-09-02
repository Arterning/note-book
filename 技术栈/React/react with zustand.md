

Create store

```js
import create from 'zustand';

interface Todo {
    id: number;
    text: string;
    done: boolean;
}

interface TodoStore {
    todos: Todo[];
    addTodo: (text: string) => void;
    toggleTodo: (id: number) => void;
}

/**
 * set 函数是 Zustand 提供的用于更新状态的方法
 * 有点类似React中的setState(state => state + 1)
 * 它接受一个回调函数，这个回调函数返回一个描述状态变化的对象。使用这个函数，你可以实现状态的更新和管理。
 */
export const useTodoStore = create<TodoStore>((set) => ({
    todos: [],
    addTodo: (text) => set((state) => ({todos: [...state.todos, {id: Date.now(), text, done: false}]})),
    toggleTodo: (id) => set((state) => ({
        todos: state.todos.map(todo => todo.id === id ? {
            ...todo,
            done: !todo.done
        } : todo)
    })),
}));

```



Use the store

```js
import { useTodoStore } from './store';

function TodoApp() {
    const todos = useTodoStore(state => state.todos);
    const addTodo = useTodoStore(state => state.addTodo);
    const toggleTodo = useTodoStore(state => state.toggleTodo);

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
}

export default TodoApp;

```















