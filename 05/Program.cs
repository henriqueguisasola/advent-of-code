using System;

namespace MyApplication
{
    class Program
    {
        static void Main(string[] args)
        {
            string[] rules = Input.Rules.Split('\n');
            string[] updates = Input.Updates.Split('\n');
            int counter = 0;

            Dictionary<string, bool> myDictionary = new Dictionary<string, bool>();
            foreach (string rule in rules)
            {
                myDictionary[rule] = true;
            }

            foreach (string update in updates)
            {
                string[] updateSplit = update.Split(',');
                Array.Reverse(updateSplit);

                bool breakingRule = false;
                for (int i = 0; i < updateSplit.Length; i++)
                {
                    for (int j = i + 1; j < updateSplit.Length; j++)
                    {
                        if (myDictionary.ContainsKey(updateSplit[i]+"|"+updateSplit[j])) {
                            breakingRule = true;
                            string temp = updateSplit[i];
                            updateSplit[i] = updateSplit[j];
                            updateSplit[j] = temp;
                        }
                    }
                }

                if (breakingRule) {
                   counter += int.Parse(updateSplit[(updateSplit.Length - 1) / 2]);
                }
            }

            Console.WriteLine(counter);
        }
    }
}
