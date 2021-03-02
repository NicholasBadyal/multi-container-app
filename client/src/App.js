import logo from './logo.svg';
import './App.css';
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import OtherPage from './otherpage';
import Fib from './fib';

function App() {
  return (
    <Router>
      <div className="App">
        <header className="App-header">
        <div className="App-subheader">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Fibonacci Calculator</h1>
          </div>
          <div className="App-Links">
          <Link className="Link" to="/">Home</Link>
          <Link className="Link" to="/otherpage">Other Page</Link>
          </div>
        </header>
        <div>
          <Route exact path="/" component={Fib} />
          <Route path="/otherpage" component={OtherPage} />
        </div>
      </div>
    </Router>
  );
}

export default App;
