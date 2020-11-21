import React from 'react';
import Header from './components/Header'
import Content from './pages/Content'
import { BrowserRouter as Router, Route, Redirect } from 'react-router-dom'
import SignIn from './pages/SignIn'


function App() {
  const uid = localStorage.getItem('uid')

  return (
    <div >
      <Router>
        {uid === null &&
          <Redirect to="/login" />
        }
        {uid != null &&
          <Header />
        }

        <Route path='/login'><SignIn /></Route>
      
        <Route exact path='/'><Content /></Route>
      </Router>
    </div>
  );
}

export default App;
