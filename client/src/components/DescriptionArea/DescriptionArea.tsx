import React from "react";

import styles from "./DescriptionArea.module.scss";

type DescriptionAreaProps = {
  editMode: boolean,
  description?: string
};

const MIN_TEXTAREA_HEIGHT = 200;

const DescriptionArea: React.FC<DescriptionAreaProps> = ({
  editMode, description=""
}) => {
  const textareaRef = React.useRef<null | HTMLTextAreaElement>(null);
  const [value, setValue] = React.useState("");

  const onChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => setValue(e.target.value);

  React.useLayoutEffect(() => {
    if (textareaRef.current) {
      textareaRef.current.style.height = "inherit";
    
      textareaRef.current.style.height = `${Math.max(
        textareaRef.current?.scrollHeight,
        MIN_TEXTAREA_HEIGHT
      )}px`;
    }
  }, [value]);

  return (
    <textarea 
      onChange={onChange}
      ref={textareaRef}
      placeholder={"Ваше описание..."}
      className={editMode ? styles.edit : styles.eval}
      readOnly={!editMode}
    >
      {description}
    </textarea>
  );
};

export default DescriptionArea;