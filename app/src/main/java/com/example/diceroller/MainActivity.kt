package com.example.diceroller

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.widget.Button
import android.widget.TextView

/**
 * このアクティビティはサイコロを振り，出た目を画面上に表示します．
 */
class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        val rollButton: Button = findViewById(R.id.button)
        rollButton.setOnClickListener {
            rollDice()
            // val resultTextView: TextView = findViewById(R.id.textView)
            // resultTextView.text = "6"
            // val toast = Toast.makeText(this, "Dice Rolled!!", Toast.LENGTH_SHORT)
            // toast.show()
            // Toast.makeText(this, "Dice Rolled!!", Toast.LENGTH_SHORT).show() でも表現できる
        }
    }

    private fun rollDice() {
        // 1つ目に6面のサイコロを作成し，振ります．
        val firstDice = Dice(6)
        val firstDiceRoll = firstDice.roll()

        // 画面上にある数字を出た目に更新します．
        val firstResultTextView: TextView = findViewById(R.id.textView1)
        firstResultTextView.text = firstDiceRoll.toString()

        // 2つ目に6面のサイコロを作成し，振ります．
        val secondDice = Dice(6)
        val secondDiceRoll = secondDice.roll()

        // 画面上にある数字を出た目に更新します．
        val secondResultTextView: TextView = findViewById(R.id.textView2)
        secondResultTextView.text = secondDiceRoll.toString()

    }
}

/**
 * サイコロを振り出た目を更新します
 */
class Dice(private val numSides: Int) {

    fun roll(): Int {
        return (1..numSides).random()
    }
}