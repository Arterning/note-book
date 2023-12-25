
## Implement Theme Swith

in the styles/global.css, we define the global css
```css
/* Light Theme */
:root {
  --background-color: white;
  --link-color: darkred;
  --text-color: black;
}

body {
  font-family: Arial, Helvetica, sans-serif;
  background-color: var(--background-color);
  color: var(--text-color);
}

a {
  color: var(--link-color);
  text-decoration: none;
}

a:focus, a:hover {
  text-decoration: underline;
}

```


first see the NavBar component

```jsx
import Link from 'next/link';
import ThemeSwitch from './ThemeSwitch';

function NavBar() {
  return (
    <nav>
      <ul>
        <li>
          <Link href="/">
            Home
          </Link>
        </li>
        <li>
          <Link href="/about">
            About
          </Link>
        </li>
        <li>
          <Link href="/test">
            Test
          </Link>
        </li>
      </ul>
      <ThemeSwitch />
      <style jsx>{`
        nav {
          display: flex;
          justify-content: space-between;
        }
        ul {
          list-style-type: none;
          padding: 0;
        }
        li {
          display: inline;
        }
        li:not(:first-child) {
          margin-left: 0.75rem;
        }
      `}</style>
    </nav>
  )
}

export default NavBar;

```

see the ThemeSwitch 

```jsx
import { useState } from 'react';
import DarkTheme from './DarkTheme';

function loadDarkMode() {
  if (typeof localStorage === 'undefined') {
    return false;
  }
  const value = localStorage.getItem('darkMode');
  return (value === null) ? false : JSON.parse(value);
}

function ThemeSwitch() {
  const [darkMode, setDarkMode] = useState(loadDarkMode);

  const handleClick = () => {
    localStorage.setItem('darkMode', JSON.stringify(!darkMode));
    setDarkMode(!darkMode);
  };

  console.log('[ThemeSwitch] darkMode:', darkMode);
  const text = darkMode ? 'Light Mode' : 'Dark Mode';
  return (
    <>
      <button onClick={handleClick} suppressHydrationWarning>
        {text}
      </button>
      <style jsx>{`
        button {
          background: none;
          border: none;
          color: inherit;
          cursor: pointer;
        }
      `}</style>
      {darkMode && <DarkTheme />}
    </>
  );
}

export default ThemeSwitch;

```

DarkTheme component is just a style piece

React Component usually include: style, js, dom

but we can create special component that only contains style or only contains js or only have dom element.

```jsx
function DarkTheme() {
  return (
    <style jsx global>{`
      :root {
        --background-color: rgb(14, 14, 14);
        --link-color: rgb(234, 207, 3);
        --text-color: rgb(230, 230, 230);
      }
    `}</style>
  );
}

export default DarkTheme;

```