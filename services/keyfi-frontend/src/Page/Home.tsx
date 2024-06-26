import './Home.css'
import Sparkle from '../assets/sparkle.png'

export default function Home() {


  return(
    <>
      <div className='home-page-container'>
        <p className='home-title'>Real estate made <span className='gradient-text'>simple</span></p>
        <p className='home-description'>Buy, sell, tour virtually, get approved 
          <br /> 
          for loans: all powered by AI 
          <img src={Sparkle} alt="" className="sparkle" />
        </p>
        <div className='input-icon'>
          <input className='home-search' type='search' placeholder="Find your newest property" />
        </div>
      </div>
    </>
  );
};