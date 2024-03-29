import React from "react";
import Square from "./Square";

export default class Board extends React.Component<
  { squares: Array<null | string>, onClick: (i: number) => void }, {}> {

  renderSquare(i: number): JSX.Element {
    return <Square 
      value={this.props.squares[i]}
      onClick={() => this.props.onClick(i)}
    />;
  }

  render(): JSX.Element {
    return (
      <div>
        <div className="board-row">
          {this.renderSquare(0)}
          {this.renderSquare(1)}
          {this.renderSquare(2)}
        </div>
        <div className="board-row">
          {this.renderSquare(3)}
          {this.renderSquare(4)}
          {this.renderSquare(5)}
        </div>
        <div className="board-row">
          {this.renderSquare(6)}
          {this.renderSquare(7)}
          {this.renderSquare(8)}
        </div>
      </div>
    );
  }
}