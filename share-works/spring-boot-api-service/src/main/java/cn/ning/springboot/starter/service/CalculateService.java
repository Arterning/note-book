package cn.ning.springboot.starter.service;


import cn.ning.springboot.starter.entity.fundaccount.Fund;
import org.springframework.stereotype.Service;

@Service
public class CalculateService {


    public void calculate(Fund fund) {

        fund.getAccounts().forEach(account -> {
            //calculate
        });

    }
}
