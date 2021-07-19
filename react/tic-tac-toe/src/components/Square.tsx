import React from "react";

export default function Square(props: {value: string|number|null, onClick: () => void}) {
  return (
    <button
      className="square"
      onClick={ props.onClick }>
      { props.value }
    </button>
  )
}