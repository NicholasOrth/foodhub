import React, {useState} from "react";
import logo from './logo.svg';
import './App.css';
import{Login} from "./Login";
import{Register} from "./Register";


function App() {
  const [state, setState] = useState('Login');
  const toggleForm = (formName) => {
    setState(formName);
  }
  return (
    <div className="App">{
      state =="Login" ? <Login onFormSwitch={toggleForm}/>: <Register onFormSwitch={toggleForm} />
    }
    </div>
  );
}

export default App;
