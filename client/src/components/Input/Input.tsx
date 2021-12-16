import React from "react";

import styles from "./Input.module.scss";

type InputProps = {
  placeholder: string,
  type?: string
}

const Input: React.FC<InputProps> = ({ placeholder, type }) => {
  return (
    <input className={styles.input} placeholder={placeholder} type={type} />
  );
};

export default Input;
