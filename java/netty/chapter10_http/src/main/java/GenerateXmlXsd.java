import java.io.IOException;

import org.jibx.binding.generator.BindGen;
import org.jibx.runtime.JiBXException;

/**
 * @author: CYH
 * @date: 2018/9/22 0022 20:08
 */
public class GenerateXmlXsd {

    public static void main(String[] args) throws JiBXException, IOException {
        BindGen.main(new String[] {"com.cyh.xml.pojo.Order"});
    }
}
