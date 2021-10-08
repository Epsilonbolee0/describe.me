import React from "react";

import styles from "./ActionTile.module.scss";

import AcceptIcon from "components/AcceptIcon";
import ActionButton from "components/ActionButton";
import ArrowLeftIcon from "components/ArrowLeftIcon";
import ArrowRightIcon from "components/ArrowRightIcon";
import CodeArea from "components/CodeArea";
import DislikeIcon from "components/DislikeIcon";
import DescriptionArea from "components/DescriptionArea";
import LikeIcon from "components/LikeIcon";
import MoveButton from "components/MoveButton";

type ActionTileProps = {
  evalMode: boolean
};

const codeTemp = `#include <iostream>

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
}`;

const descrTemp = `Функция представляет собой цикл for, который делает прекрасные вещи,\
 что-то к чему то прибавляет, сообщение выводит, а возвращает забавную фразу`;


const ActionTile: React.FC<ActionTileProps> = ({evalMode}) => {
  return (
    <div className={styles.action_tile}>
      <div className={styles.action_header}>
        <MoveButton children1={<ArrowLeftIcon />} children2={"Назад"}></MoveButton>

        <h2 className={styles.lang_header}>Язык: C</h2>
        
        <MoveButton children1={"Вперёд"} children2={<ArrowRightIcon />}></MoveButton>
      </div>

      <CodeArea language={"c"} codeText={codeTemp} />

      <div className={styles.action_wrapper}>
        <DescriptionArea editMode={evalMode} />

        <div className={styles.button_wrapper}>
          <ActionButton acceptMode={true}>
            {!evalMode ? <LikeIcon /> : <AcceptIcon />}
          </ActionButton>

          {!evalMode && <ActionButton acceptMode={false}><DislikeIcon /></ActionButton>}
        </div>
      </div>
    </div>
  );
}

export default ActionTile;