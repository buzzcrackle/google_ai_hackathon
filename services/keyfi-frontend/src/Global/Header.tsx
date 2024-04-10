import './Header.css';
import logo from '../assets/key.png';

export default function Header() {
  return(
    <header>
      <nav className="navbar-container">
        <ul>
          <li><a href="/"><img src={logo} alt="Website logo" className="logo-image" /></a></li>
          <li><a href="/">Home</a></li>
          <li><a href="/chat">AI Chatbot</a></li>
          <li><a href="/map">Map</a></li>
        </ul>
      </nav>
    </header>
  );
}