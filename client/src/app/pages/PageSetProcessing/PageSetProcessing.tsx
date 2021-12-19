import React from "react";

import styles from "./PageSetProcessing.module.scss";
import ActionTile from "components/ActionTile";
import Header from "components/Header";
import { ActionType } from "utils/types";

type PageSetProcessingProps = {

}

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


const PageSetProcessing: React.FC<PageSetProcessingProps> = () => {
  return (
    <div>
      <div className={styles.tile_container}>
        <ActionTile mode={ActionType.eval} lang={"C"} code={codeTemp} description={descrTemp}/>
      </div>
    </div>
  ); 
};

export default PageSetProcessing;
