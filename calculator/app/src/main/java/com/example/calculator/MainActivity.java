package com.example.calculator;

import android.app.Activity;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;

public class MainActivity extends Activity {

    private TextView display;

    private double firstNumber = 0;
    private String operator = "";
    private boolean newNumber = true;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        setContentView(R.layout.activity_main);

        display = findViewById(R.id.display);

        setupNumberButtons();
        setupOperatorButtons();
    }

    private void setupNumberButtons() {

        int[] numberIds = {
                R.id.zero,
                R.id.one,
                R.id.two,
                R.id.three,
                R.id.four,
                R.id.five,
                R.id.six,
                R.id.seven,
                R.id.eight,
                R.id.nine
        };

        for (int id : numberIds) {

            Button button = findViewById(id);

            button.setOnClickListener(v -> {

                String number = ((Button) v).getText().toString();

                if (newNumber || display.getText().toString().equals("0")) {
                    display.setText(number);
                    newNumber = false;
                } else {
                    display.append(number);
                }
            });
        }

        findViewById(R.id.decimal).setOnClickListener(v -> {

            String current = display.getText().toString();

            if (newNumber) {
                display.setText("0.");
                newNumber = false;
            } else if (!current.contains(".")) {
                display.append(".");
            }
        });
    }

    private void setupOperatorButtons() {

        findViewById(R.id.add).setOnClickListener(v -> setOperator("+"));

        findViewById(R.id.subtract).setOnClickListener(v -> setOperator("-"));

        findViewById(R.id.multiply).setOnClickListener(v -> setOperator("*"));

        findViewById(R.id.divide).setOnClickListener(v -> setOperator("/"));

        findViewById(R.id.equal).setOnClickListener(v -> calculate());

        findViewById(R.id.clear).setOnClickListener(v -> clear());

        findViewById(R.id.delete).setOnClickListener(v -> delete());
    }

    private void setOperator(String selectedOperator) {

        firstNumber = Double.parseDouble(display.getText().toString());

        operator = selectedOperator;

        newNumber = true;
    }

    private void calculate() {

        if (operator.isEmpty()) {
            return;
        }

        double secondNumber;

        try {
            secondNumber = Double.parseDouble(display.getText().toString());
        } catch (Exception e) {
            display.setText("Error");
            return;
        }

        double result = 0;

        switch (operator) {

            case "+":
                result = firstNumber + secondNumber;
                break;

            case "-":
                result = firstNumber - secondNumber;
                break;

            case "*":
                result = firstNumber * secondNumber;
                break;

            case "/":

                if (secondNumber == 0) {
                    display.setText("Cannot divide by 0");
                    operator = "";
                    newNumber = true;
                    return;
                }

                result = firstNumber / secondNumber;
                break;
        }

        if (result == (long) result) {
            display.setText(String.valueOf((long) result));
        } else {
            display.setText(String.valueOf(result));
        }

        operator = "";
        newNumber = true;
    }

    private void clear() {

        display.setText("0");

        firstNumber = 0;

        operator = "";

        newNumber = true;
    }

    private void delete() {

        String current = display.getText().toString();

        if (current.length() <= 1) {
            display.setText("0");
            newNumber = true;
            return;
        }

        display.setText(current.substring(0, current.length() - 1));
    }
}
