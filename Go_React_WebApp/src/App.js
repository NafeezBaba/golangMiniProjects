import React, { Component } from 'react';
// import logo from './logo.svg';
import './App.css';
// import Form from './container/form'
import Appbar from './components/appbar'

class App extends Component {
  render() {
    return (
      <div className="App">
       
      <Appbar/>
      {/* <Form/> */}
      </div>
    );
  }
}

export default App;
