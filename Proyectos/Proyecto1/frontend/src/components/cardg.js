import React from 'react';
import '../css/cardG.css';

const CardG = ({children}) => {
    return (
        <div className="card2 shadow">
            <div className="card2__content">
                {children}
            </div>
        </div>
    );
};

export default CardG;
