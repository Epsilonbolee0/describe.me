import React from "react";

import styles from "./ActionTile.module.scss";

import ArrowLeftIcon from "components/ArrowLeftIcon";
import ArrowRightIcon from "components/ArrowRightIcon";
import CodeArea from "components/CodeArea";
import MoveButton from "components/MoveButton";

type ActionTileProps = {
}

const ActionTile: React.FC<ActionTileProps> = () => {
  return (
    <div className={styles.action_tile}>
      <div className={styles.action_header}>
        <MoveButton children1={<ArrowLeftIcon />} children2={"Назад"}></MoveButton>

        <h2 className={styles.lang_header}>Язык: Go-lang</h2>

        <MoveButton children1={"Вперёд"} children2={<ArrowRightIcon />}></MoveButton>
      </div>
      <CodeArea language={"c"} codeText={
        `#include <iostream>

int main(void) 
{ 
    int index = 1; 
    int array[] = {1, 2, 3, 4, 5, 6}; 
  
    for (index; index < 10; index++) 
    { 
        bubbleSort(&array, index); 
        printf(“Success”); 
    }

    return "I like write shit"
}`
      }/>
    </div>
  );
}

export default ActionTile;