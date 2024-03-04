import React from 'react';
import '../css/card.css';

const Card = ({title, children}) => {
    return (
        <div>
            <div className="card">
                <div className="tools">
                    <div className="circle">
                        <span className="red box"></span>
                    </div>
                    <div className="circle">
                        <span className="yellow box"></span>
                    </div>
                    <div className="circle">
                        <span className="green box"></span>
                    </div>
                    <div className="titulo">
                        {title}
                    </div>
                </div>
                <div className="card__content">
                    {children}
                </div>
            </div>
        </div>
    );
};

export default Card;
