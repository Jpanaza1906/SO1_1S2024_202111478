import React from "react";
import "../css/combobox.css";

const ComboBox = ({ options, onChange, placeholder }) => {
    return (
        <div className="select-wrapper">
            <select onChange={onChange}>
                {placeholder && <option value="">{placeholder}</option>}
                {options.map((option, index) => (
                    <option key={index} value={option}>{option}</option>
                ))}
            </select>
        </div>
    );
};

export default ComboBox;