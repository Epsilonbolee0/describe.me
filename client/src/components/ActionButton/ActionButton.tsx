import React from "react";

import styles from "./ActionButton.module.scss";

type ActionButtonProps = {
  acceptMode: boolean,
  children: React.ReactNode
};

const ActionButton: React.FC<ActionButtonProps> = ({acceptMode, children}) => {
  return (
    <button 
      className={acceptMode ? styles.accept : styles.decline}
    >
      {children}
    </button>
  );
};

export default ActionButton;