import React from "react";

import styles from "./CodeArea.module.scss";

type CodeAreaProps = {
  codeText: string,
  language: string
}

const CodeArea: React.FC<CodeAreaProps> = ({codeText, language}) => {
  return (
    <pre className={styles.code_area}>
      <code data-language={language}>
        {codeText}
      </code>
    </pre> 
  );
};

export default CodeArea;