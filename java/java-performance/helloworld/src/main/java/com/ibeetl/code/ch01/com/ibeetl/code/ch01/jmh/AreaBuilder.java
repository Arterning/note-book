package com.ibeetl.code.ch01.com.ibeetl.code.ch01.jmh;

import java.util.ArrayList;
import java.util.List;
import java.util.function.Supplier;

public class AreaBuilder {

	/**
	 * 模拟测试数据,建议使用较大size以验证效果
	 * @param size
	 * @return
	 */
	public static List<Org> initSize(int size){
		Org org = new Org();
		org.setProvinceId(21);
		org.setName("北京");
		ArrayList<Org> list = new ArrayList();
		for(int i=0;i<size;i++){
			list.add(org);
		}
		return list;
	}
	/**
	 * 取出省id,拼接字符串，此代码有7处可能的优化地方
	 * @param orgs
	 * @return
	 */
	public static String buildProvince(List<Org> orgs){
		StringBuilder sb = new StringBuilder();
		for(Org org:orgs){
			if(sb.length()!=0){
				sb.append(",");
			}
			sb.append(org.getProvinceId());
		}
		return sb.toString();
	}

	/*********************** 一下是优化方法    *******************/
	public static String buildProvinceWithInitCapacity(List<Org> orgs){
		//假设省Id不超过2位，初始化一个大小
		StringBuilder sb = new StringBuilder(orgs.size()*3);
		for(Org org:orgs){
			if(sb.length()!=0){
				sb.append(",");
			}
			sb.append(org.getProvinceId());
		}
		return sb.toString();
	}


	public static String buildProvinceWithStringBuffer(List<Org> orgs){
		//StringBuffer
		StringBuffer sb = new StringBuffer();
		for(Org org:orgs){
			if(sb.length()!=0){
				sb.append(",");
			}
			sb.append(org.getProvinceId());
		}
		return sb.toString();
	}

	public static String buildProvinceWithEmptyCheck(List<Org> orgs){
		//检测是否为空
		if(orgs.isEmpty()){
			return "";
		}
		StringBuilder sb = new StringBuilder();
		for(Org org:orgs){
			if(sb.length()!=0){
				sb.append(",");
			}
			sb.append(org.getProvinceId());
		}
		return sb.toString();
	}


	public static String buildProvinceWithSizeCheck(List<Org> orgs){
		//如果只包含一个元素
		if(orgs.size()==1){
			return String.valueOf(orgs.get(0).getProvinceId());
		}
		StringBuilder sb = new StringBuilder();
		for(Org org:orgs){
			if(sb.length()!=0){
				sb.append(",");
			}
			sb.append(org.getProvinceId());
		}
		return sb.toString();
	}


	public static String buildProvinceWithSubString(List<Org> orgs){
		StringBuilder sb = new StringBuilder();
		for(Org org:orgs){
			sb.append(org.getProvinceId()).append(",");
		}
		//去掉最后一个逗号
		sb.setLength(sb.length()-1);
		return sb.toString();
//		return sb.substring(0,sb.length());

	}

	public static String buildProvinceWithChar(List<Org> orgs){
		StringBuilder sb = new StringBuilder();
		for(Org org:orgs){
			//使用char 代替String
			sb.append(org.getProvinceId()).append(',');
		}
		sb.setLength(sb.length()-1);
		return sb.toString();
	}

	/************ 比较冷门的优化方法 **************/

	public static String buildProvinceWithCacheData(List<Org> orgs){
		StringBuilder sb = new StringBuilder();
		for(Org org:orgs){
			//使用预先初始化好的chars
			char[] chars = strProvince[org.getProvinceId()];
			sb.append(chars).append(',');
		}
		sb.setLength(sb.length()-1);
		return sb.toString();
	}

	static char[][]  strProvince;
	static{
		strProvince = new char[31][];
		for(int i=0;i<31;i++){
			strProvince[i] = String.valueOf(i).toCharArray();
		}
	}


	public static String buildProvinceWithCache(List<Org> orgs){
		StringBuilder sb = local.get();
		for(Org org:orgs){
			if(sb.length()!=0){
				sb.append(",");
			}
			sb.append(org.getProvinceId());
		}
		String str =  sb.toString();
		sb.setLength(0);
		return str;
	}

	static ThreadLocal<StringBuilder> local = ThreadLocal.withInitial((Supplier) () -> new StringBuilder());


	/**
	 * 如果预先知道容量，比如能整除4，可以减少循环判断
	 * @param orgs
	 * @return
	 */
	public static String buildProvinceWithLoop(List<Org> orgs){
		StringBuilder sb = new StringBuilder();

		int count = orgs.size()/4;
		int i=0;
		for(;i<count;i++){
			sb.append(orgs.get(i).getProvinceId()).append(",");
			sb.append(orgs.get(i+1).getProvinceId()).append(",");
			sb.append(orgs.get(i+2).getProvinceId()).append(",");
			sb.append(orgs.get(i+3).getProvinceId()).append(",");
		}

		//去掉最后一个逗号
		sb.setLength(sb.length()-1);
		return sb.toString();
	}





}
