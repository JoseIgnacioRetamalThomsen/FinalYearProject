import React from 'react';
import HelloWorld from './HelloWorld';
import Login from './components/Login';

//HelloWorld.sayHello("Hello from native-react");

const App: () => React$Node = () => {
  return (
    <>
      <Login/>
    </>
  );
};

export default App;
