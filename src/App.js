import React from 'react';
import Header from './components/Header'
import Content from './pages/Content'
import Post from './pages/Post'
import { BrowserRouter as Router, Route, Redirect } from 'react-router-dom'
import SignIn from './pages/SignIn'
function App() {
  const uid = localStorage.getItem('uid')

  return (
    <div >
      <Router>

        <Header />



        <Route path='/login'><SignIn /></Route>
        <Route path='/post'><Post /></Route>
        <Route exact path='/'><Content /></Route>

      </Router>
    </div>
  );
}

export default App;
