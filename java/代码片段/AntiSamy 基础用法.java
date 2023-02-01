import org.owasp.validator.html.AntiSamy;
import org.owasp.validator.html.CleanResults;
import org.owasp.validator.html.Policy;
import org.owasp.validator.html.PolicyException;

import java.io.File;
import java.io.IOException;
/**
 * From ChatGPT
 * 
 * In this example, the AntiSamy policy file is loaded from a file named "antisamy.xml". The scan method of the AntiSamy class is then used to clean the user-generated input dirtyInput. The cleaned HTML is stored in the CleanResults object cr, which can be used to access the cleaned HTML, a list of error messages, and a list of the malicious input that was removed. In this example, the cleaned HTML is extracted and printed to the console.
 */
public class AntiSamyExample {

    public static void main(String[] args) {
        try {
            // Load the AntiSamy policy file
            Policy policy = Policy.getInstance(new File("antisamy.xml"));

            // Create an instance of the AntiSamy class
            AntiSamy antiSamy = new AntiSamy();

            // Get the untrusted input from the user
            String dirtyInput = "<script>alert('XSS');</script><p>Hello World!</p>";

            // Scan the input using the AntiSamy class
            CleanResults cr = antiSamy.scan(dirtyInput, policy);

            // Get the clean output
            String cleanOutput = cr.getCleanHTML();

            // Print the clean output
            System.out.println("Clean Output: " + cleanOutput);

        } catch (PolicyException | IOException e) {
            e.printStackTrace();
        }
    }
}
